<script>
import "./app.css";
import Header from "./components/Header.svelte";

let userIdentifier;
let email;
let password;


async function handleLogin(e) {
  e.preventDefault();

  const userCredentials = {
    userIdentifier,
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

<Header />

<main>
  <section class="login login__container w-full flex flex-col items-center">
    <div class="login__column column w-5/6 md:4/5 lg:w-1/3">
      <form on:submit={handleLogin} id="login" class="login-form flex flex-col px-20 py-2 text-center border border-gray-300">
        <h1 class="login-form__logo mt-9 mb-3">Svelte.IG</h1>
        <label for="userIdentifier"></label>
        <input 
        class="login-form__user-input h-10 mb-2 px-2 text-sm rounded border border-gray-300"
        type="text"
        id="user-identifier"
        name="userIdentifier"
        placeholder="Phone number, username, or email"
        bind:value={userIdentifier}
        required
        >
        <input
        class="login-form__password-input h-10 mb-4 px-2 text-sm rounded border border-gray-300"
        type="password"
        id="password"
        name="password"
        placeholder="Password"
        bind:value={password} 
        required
        >
        <input 
        class="login-form__submi-inputt h-10 mb-2 px-2 text-sm rounded-lg bg-blue-500 text-white"
        type="submit" 
        value="Log In"
        >
      </form>
    </div>
  </section>
</main>

<style>

  main {
    height: 100%;
    min-height: 100vh;
    max-height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .login__container {
    height: 100%;
  }

  .login-form__logo {
    font-family: 'Borel', cursive;
    font-size: 32px;
  }

</style>
