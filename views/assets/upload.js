"use strict";

document.addEventListener("alpine:init", () => {
  Alpine.data("global", () => ({
    isLoading: false,
    toastContent: { isOpen: false },
    files: [],
    progress: 0,
    session: "",
    toggleToast(isOpen, data) {
      this.toastContent = {
        isOpen,
        ...data,
      };
    },
    handleFileChange(event) {
      this.files = this.files.concat(Array.from(event.target.files));
    },
    handleDrop(event) {
      event.preventDefault();
      this.files = this.files.concat(Array.from(event.dataTransfer.files));
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
      this.files.forEach((file) => formData.append("files", file));

      const path = "/api/v1/upload" + "?s=" + this.session;

      fetch(path, {
        method: "POST",
        body: formData,
        headers: {
          "X-Requested-With": "XMLHttpRequest", // To identify AJAX request on the server
        },
        onUploadProgress: (progressEvent) => {
          this.progress = Math.round(
            (progressEvent.loaded / progressEvent.total) * 100
          );
        },
      })
        .then((response) => {
          return response.json();
        })
        .then(() => {
          this.toggleToast(true, {
            type: "success",
            content: `${this.files.length} files were transfered succesfully!`,
          });
          this.files = [];
          this.progress = 0;
          setTimeout(() => {
            this.toggleToast(false);
          }, 6000);
        })
        .catch((error) => {
          console.error("Error uploading files:", error);
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  }));
});
