# Web Client for the Hello Server

Ensure that your web API server in the `helloserver/` directory is running and ready for reqeusts. Then create a new web app in this directory that meets the following requirements:

- Your page must have an `<input>` element into which the user can type their name.
- Your page must have a `<button>` element the user can click/tap to submit that data. You may also let the user trigger a submit by hitting the `Enter` key while focus is in the `<input>` element (hint: use a `<form>` element and catch the form's `submit` event).
- When the user submits the data, do two things:
	- Use AJAX to make a `GET /?name={name-entered}` request to your web server, where `{name-entered}` is replaced by the value of your `<input>` element. Render the text you get back into some element on the page so that the user can see the welcome message.
	- Use an `<img>` element to display the identicon for the name in your `<input>` element. Remember that the `/identicon/{name-entered}` resource path of your web server will return the identicon PNG image bytes for the name you provide.
