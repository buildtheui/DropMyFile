"use strict";

const TOAST_TIMEOUT = 6000;
let WSObserver;

document.addEventListener("alpine:init", () => {
  Alpine.data("global", () => ({
    search: "",
    isLoading: false,
    toastContent: { isOpen: false },
    filesToUpload: [],
    filesList: [],
    session: "",
    initData(session) {
      this.listenFileChanges(session);
      return session;
    },

    get filteredFiles() {
      return this.filesList.filter((file) =>
        file.fileName.toLowerCase().includes(this.search.toLowerCase())
      );
    },
    listenFileChanges(session) {
      const filesWSLink = `ws://${window.location.host}/ws/files?s=${session}`;
      WSObserver = new window.WSObservable({
        wsLink: filesWSLink,
      });

      const subscriber = WSObserver.subscribe({
        next: (event, data) => {
          if (event === "files_modified") {
            console.log(data);
            this.filesList = data ?? [];
          }
        },
        // TODO: handle this error case
        error: () => {
          this.toggleToast(true, {
            type: "error",
            content: "Error listening changes from files",
          });
        },
        // TODO: handle this complete case
        complete: () => {
          subscriber.unsubscribe();
          this.toggleToast(
            true,
            {
              type: "error",
              content:
                "Connection close: To see changes try restarting DropMyFile",
            },
            false
          );
        },
      });
    },
    toggleToast(isOpen, data, closeWithTimeout = true) {
      this.toastContent = {
        isOpen,
        ...data,
      };

      closeWithTimeout &&
        this.toastContent.isOpen &&
        setTimeout(() => {
          this.toastContent = { isOpen: false };
        }, TOAST_TIMEOUT);
    },
    handleFileChange(event) {
      this.filesToUpload = this.filesToUpload.concat(
        Array.from(event.target.files)
      );
    },
    handleDrop(event) {
      event.preventDefault();
      this.filesToUpload = this.filesToUpload.concat(
        Array.from(event.dataTransfer.files)
      );
    },
    formatBytes(bytes, decimals = 2) {
      if (bytes === 0) return "0 Bytes";
      const k = 1024;
      const dm = decimals < 0 ? 0 : decimals;
      const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
      const i = Math.floor(Math.log(bytes) / Math.log(k));
      return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
    },
    uploadFiles() {
      if (this.isLoading) return;

      this.isLoading = true;

      const formData = new FormData();
      this.filesToUpload.forEach((file) => formData.append("files", file));

      const path = "/api/v1/upload" + "?s=" + this.session;

      fetch(path, {
        method: "POST",
        body: formData,
        headers: {
          "X-Requested-With": "XMLHttpRequest", // To identify AJAX request on the server
        },
      })
        .then((response) => {
          return response.json();
        })
        .then(() => {
          this.toggleToast(true, {
            type: "success",
            content: `${this.filesToUpload.length} files were transfered succesfully!`,
          });
          this.filesToUpload = [];
          this.$refs.fileInput.value = "";
        })
        .catch((error) => {
          this.toggleToast(true, {
            type: "error",
            content: "Error uploading: try restarting DropMyFile",
          });
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  }));
});
