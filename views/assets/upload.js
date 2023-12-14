function handleFileChange(event) {
  this.files = Array.from(event.target.files);
}

function handleDrop(event) {
  event.preventDefault();
  this.files = Array.from(event.dataTransfer.files);
}

function formatBytes(bytes, decimals = 2) {
  if (bytes === 0) return "0 Bytes";
  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
}

function uploadFiles() {
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
    .then((response) => response.json())
    .then((data) => {
      console.log("Files uploaded:", data);
      this.files = [];
      this.progress = 0;
      // Handle the server response as needed
    })
    .catch((error) => {
      console.error("Error uploading files:", error);
    });
}
