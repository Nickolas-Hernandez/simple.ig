<script>
import "./app.css"

let username;
let email;
let password;


async function handleLogin(e) {
  e.preventDefault();

  const userCredentials = {
    username,
    email,
    password
  };

  try{
    const response = await fetch('/api/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(userCredentials)
    });
    
    if (response.ok) {
      console.log('response: ', response);
      const userToken = await response.json();
      console.log(userToken);
      // Handle login success, e.g., redirect to homepage
      // window.location.replace('/'
    } else {
      // Handle login failure, e.g., show an error message to the user
      console.log('else ', response);
    }
  } catch (err) {
    console.error(err);
  }
}

</script>

<main>
  <section class="login">
    <h1>Create an Account</h1>
    <form on:submit={handleLogin} id="login">
      <label for="username">Username</label>
      <input type="text" id="username" name="username" bind:value={username} required>
      <label for="email">Email</label>
      <input type="text" id="email" name="email" bind:value={email} required>
      <label for="password">Password</label>
      <input type="password" id="password" name="password" bind:value={password} required>
      <input type="submit" value="Ping">
    </form>
  </section>
</main>

<style>

</style>
