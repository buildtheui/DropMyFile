class WSObservable {
  #ws;
  #cacheRes;

  constructor({ wsLink }) {
    this.#initWSConnection(wsLink);
  }

  subscribe(observer) {
    if (typeof observer !== "object" || observer === null) {
      throw new Error(
        "Observer must be an object with next, error, and complete methods"
      );
    }

    if (typeof observer.next !== "function") {
      throw new Error("Observer must have a next method");
    }

    if (typeof observer.error !== "function") {
      throw new Error("Observer must have an error method");
    }

    if (typeof observer.complete !== "function") {
      throw new Error("Observer must have a complete method");
    }

    const unsubscribe = this.#producer(observer);

    return {
      unsubscribe() {
        if (unsubscribe && typeof unsubscribe === "function") {
          unsubscribe();
        }
      },
    };
  }

  #initWSConnection(wsLink) {
    this.#ws = new WebSocket(wsLink);
  }

  #producer(observer) {
    const socket = this.#ws;

    if (socket == null) {
      throw new Error("Can not observe an open web socket");
    }

    const openListener = (_event) => {
      // future open listener logic
    };

    const messageListener = (event) => {
      if (!event.data) return;

      if (event.data == null || !this.#isDataNew(event.data)) {
        return;
      }

      const data = JSON.parse(event.data);
      observer.next(data.eventName, data.payload);
    };

    const closeListener = () => {
      observer.complete();
    };

    const errorListener = (event) => {
      observer.error(event);
    };

    // Listen for open connection
    socket.addEventListener("open", openListener);

    // Listen for messages
    socket.addEventListener("message", messageListener);

    // Connection closed
    socket.addEventListener("close", closeListener);

    // Connection error
    socket.addEventListener("error", errorListener);

    return () => {
      socket.removeEventListener("open", openListener);
      socket.removeEventListener("message", messageListener);
      socket.removeEventListener("close", closeListener);
      socket.removeEventListener("error", errorListener);

      socket.close();
    };
  }

  #isDataNew(newData) {
    if (newData !== this.#cacheRes) {
      this.#cacheRes = newData;
      return true;
    }

    return false;
  }
}

window.WSObservable = WSObservable;
