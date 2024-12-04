// Write EventListener class

class EventEmitter {
    constructor(){
      this.events = {}
    }
    on(eventName, handler) {
      this.events[eventName] = {
        func: handler
      }
    }
    once(eventName, handler) {
      this.events[eventName] = {
        func: handler,
        once: true
      }
    }
    removeListener(eventName) {
      delete this.events[eventName]
    }
    removeAllListeners(eventName) {
      this.events = {}
    }
    
    emit(eventName, ...args) {
      if (this.events[eventName]) {
        this.events[eventName].func(...args);
        
        if (this.events[eventName].once) 
          this.removeListener(eventName)
      }
  
    }
  
  }
  
const eventEmitter = new EventEmitter();

eventEmitter.on('start', number => {
    console.log('started', number);
});

eventEmitter.once('wow', number => {
    console.log('wow triggered', number);
});

eventEmitter.emit('start', 23);
eventEmitter.emit('wow', 1);
eventEmitter.emit('wow', 2);
eventEmitter.removeListener('start')
eventEmitter.removeAllListeners()
  