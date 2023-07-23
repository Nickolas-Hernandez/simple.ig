<script>
let username = '';
let password = '';


async function handleLogin(e) {
  e.preventDefault();
  const response = await fetch('/api/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      username: 'example_user',
      password: 'example_password'
    })
  });

  if (response.ok) {
    // Handle successful login, e.g., store authentication token in browser's local storage
    console.log('success!');
  } else {
    // Handle login failure, e.g., show an error message to the user
    console.log('fail :(');
  }
}

</script>

<main>
  <section class="login">
    <h1>Login</h1>
    <form on:submit={handleLogin} id="login">
      <input type="text" id="username" name="username" required>
      <input type="password" id="password" name="password" required>
      <input type="submit" value="login">
    </form>
  </section>

</main>

<style>

</style>
