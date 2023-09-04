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
  <section class="login w-full flex flex-col items-center">
    <div class="column w-5/6 md:4/5 lg:w-1/3">
      <form on:submit={handleLogin} id="login" class="bg-blue-100 flex flex-col px-20 text-center">
        <h1>Svelte.IG</h1>
        <input type="text" id="username" name="username" bind:value={username} required>
        <input type="text" id="email" name="email" bind:value={email} required>
        <input type="password" id="password" name="password" bind:value={password} required>
        <input type="submit" value="Ping">
      </form>
    </div>
  </section>
</main>

<style>

  .login form input {
    margin-bottom: 6px;
    border: 1px solid #ccc;
  }

</style>
