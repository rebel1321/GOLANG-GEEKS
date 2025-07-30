package observer

type Observer struct {
	Name                string
	NotificationChannel chan string
}

type Subject interface {
	Register(observer *Observer)
	Deregister(observer *Observer)
	NotifyAll(msg string)
}

type NewsAgency struct {
	Observers []*Observer
}

// Register an observer
func (a *NewsAgency) Register(observer *Observer) {
	a.Observers = append(a.Observers, observer)
}

// Deregister an observer
func (a *NewsAgency) Deregister(observer *Observer) {
	for i, o := range a.Observers {
		if o == observer {
			a.Observers = append(a.Observers[:i], a.Observers[i+1:]...)
			break
		}
	}
}

// Notify all observers
func (a *NewsAgency) NotifyAll(msg string) {
	for _, o := range a.Observers {
		notification := "Message is sent to the observer named " + o.Name + ": " + msg
		o.NotificationChannel <- notification
	}
}