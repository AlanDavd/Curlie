document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("curlForm");
  const addHeaderBtn = document.getElementById("addHeader");
  const addQueryBtn = document.getElementById("addQuery");
  const copyBtn = document.getElementById("copyCommand");
  const resultPre = document.getElementById("result");

  // Add new header input pair
  addHeaderBtn.addEventListener("click", () => {
    const headerPair = document.createElement("div");
    headerPair.className = "header-pair grid grid-cols-2 gap-2";
    headerPair.innerHTML = `
            <input type="text" class="header-key mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="Key">
            <input type="text" class="header-value mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="Value">
        `;
    document.getElementById("headers").appendChild(headerPair);
  });

  // Add new query parameter input pair
  addQueryBtn.addEventListener("click", () => {
    const queryPair = document.createElement("div");
    queryPair.className = "query-pair grid grid-cols-2 gap-2";
    queryPair.innerHTML = `
            <input type="text" class="query-key mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="Key">
            <input type="text" class="query-value mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500" placeholder="Value">
        `;
    document.getElementById("queryParams").appendChild(queryPair);
  });

  // Handle form submission
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    // Collect headers
    const headers = {};
    document.querySelectorAll(".header-pair").forEach((pair) => {
      const key = pair.querySelector(".header-key").value.trim();
      const value = pair.querySelector(".header-value").value.trim();
      if (key && value) {
        headers[key] = value;
      }
    });

    // Collect query parameters
    const queryParams = {};
    document.querySelectorAll(".query-pair").forEach((pair) => {
      const key = pair.querySelector(".query-key").value.trim();
      const value = pair.querySelector(".query-value").value.trim();
      if (key && value) {
        queryParams[key] = value;
      }
    });

    const data = {
      method: document.getElementById("method").value,
      url: document.getElementById("url").value,
      headers: headers,
      queryParams: queryParams,
      body: document.getElementById("body").value.trim(),
    };

    try {
      const response = await fetch("/api/curl", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      const result = await response.json();
      if (response.ok) {
        resultPre.textContent = result.command;
      } else {
        resultPre.textContent = `Error: ${result.error}`;
      }
    } catch (error) {
      resultPre.textContent = `Error: ${error.message}`;
    }
  });

  // Copy command to clipboard
  copyBtn.addEventListener("click", () => {
    const command = resultPre.textContent;
    if (command) {
      navigator.clipboard
        .writeText(command)
        .then(() => {
          const originalText = copyBtn.textContent;
          copyBtn.textContent = "Copied!";
          setTimeout(() => {
            copyBtn.textContent = originalText;
          }, 2000);
        })
        .catch((err) => {
          console.error("Failed to copy text: ", err);
        });
    }
  });
});
