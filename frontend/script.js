// JavaScript code for handling frontend interactions

// Example code for handling post form submission
document.querySelector('form').addEventListener('submit', function (event) {
    event.preventDefault();
    const title = document.getElementById('post-title').value;
    const content = document.getElementById('post-content').value;

    // Send the data to the backend API for creating a new post
    fetch('/api/posts', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            title: title,
            content: content,
        }),
    })
    .then(response => response.json())
    .then(data => {
        // Handle the response from the backend (e.g., show a success message)
        console.log(data);
    })
    .catch(error => {
        // Handle errors if any
        console.error('Error:', error);
    });
});
