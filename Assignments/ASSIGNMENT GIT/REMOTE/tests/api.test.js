const axios = require("axios");

// Example API test (dummy URL)
test("API returns success status", async () => {
    const response = await axios.get("https://jsonplaceholder.typicode.com/posts/1");
    expect(response.status).toBe(200);
});
