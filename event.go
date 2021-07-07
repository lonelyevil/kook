package khl

// EventHandler is the interface for objects handling event.
type EventHandler interface {
	Type() string
	Handle(*Session, *EventDataGeneral, interface{})
}

// EventHandlerProvider is the interface for objects providing event handlers.
type EventHandlerProvider interface {
	Type() string
	New() interface{}
}

var registeredSystemEventHandler = map[string]EventHandlerProvider{}

func registerSystemEventHandler(seh EventHandlerProvider) {
	if _, ok := registeredSystemEventHandler[seh.Type()]; ok {
		return
	}
	registeredSystemEventHandler[seh.Type()] = seh
	return
}

var registeredMessageEventHandler = map[string]EventHandlerProvider{}

func registerMessageEventHandler(meh EventHandlerProvider) {
	if _, ok := registeredMessageEventHandler[meh.Type()]; ok {
		return
	}
	registeredMessageEventHandler[meh.Type()] = meh
	return
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

func (s *Session) handle(t string, edg *EventDataGeneral, i interface{}) {
	for _, eh := range s.handlers[t] {
		eh.eventHandler.Handle(s, edg, i)
	}
}

func (s *Session) handleEvent(t string, edg *EventDataGeneral, i interface{}) {
	s.handlersMu.RLock()
	defer s.handlersMu.RUnlock()
	s.handle(t, edg, i)
}
