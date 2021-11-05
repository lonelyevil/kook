package khl

// EventHandler is the interface for objects handling event.
type EventHandler interface {
	Type() string
	Handle(EventContext)
}

// EventHandlerProvider is the interface for objects providing event handlers.
type EventHandlerProvider interface {
	Type() string
	New() EventContext
}

// EventContext is the interface for objects containing context for event handlers.
type EventContext interface {
	GetExtra() interface{}
	GetCommon() *EventHandlerCommonContext
}

var registeredEventHandler = map[string]EventHandlerProvider{}

func registerEventHandler(eh EventHandlerProvider) {
	if _, ok := registeredEventHandler[eh.Type()]; ok {
		return
	}
	registeredEventHandler[eh.Type()] = eh
}

type eventHandlerInstance struct {
	eventHandler EventHandler
}

func (s *Session) addEventHandler(handler EventHandler) func() {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()

	if s.handlers == nil {
		//s.log(LogTrace, "creating initial handlers")
		addCaller(s.Logger.Trace()).Msg("creating initial handlers")
		s.handlers = map[string][]*eventHandlerInstance{}
	}

	ehi := &eventHandlerInstance{eventHandler: handler}
	s.handlers[handler.Type()] = append(s.handlers[handler.Type()], ehi)

	return func() {
		s.removeEventHandler(handler.Type(), ehi)
	}
}

// AddHandler adds event handlers to session, and provides additional type check.
func (s *Session) AddHandler(h interface{}) func() {
	eh := handlerForInterface(h)
	if eh == nil {
		//s.log(LogError, "Invalid handler type, ignored")
		addCaller(s.Logger.Error()).Msg("Invalid handler type, ignored")
		return func() {
			// TODO: add remover.
		}
	}
	return s.addEventHandler(eh)
}

func (s *Session) removeEventHandler(t string, ehi *eventHandlerInstance) {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()
	handlers := s.handlers[t]
	for i := range handlers {
		if handlers[i] == ehi {
			s.handlers[t] = append(handlers[:i], handlers[i+1:]...)
		}
	}
}

func (s *Session) handle(t string, i EventContext) {
	for _, eh := range s.handlers[t] {
		if s.Sync {
			eh.eventHandler.Handle(i)
		} else {
			go eh.eventHandler.Handle(i)
		}
	}
}

func (s *Session) handleEvent(t string, edg *EventDataGeneral, i EventContext) {
	s.handlersMu.RLock()
	defer s.handlersMu.RUnlock()
	c := i.GetCommon()
	c.Common = edg
	c.Session = s
	s.handle(t, i)
}
