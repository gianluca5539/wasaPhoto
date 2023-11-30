<script>
export default {
  data() {
    return {
      loading: false,
      username: ''
    };
  },
  methods: {
    async login(e) {
      e.preventDefault();
      if (this.username.trim() !== '') {
        this.loading = true;

        // TODO get the uuid from the backend
        let uuid = 'asdassda';
        if (uuid != '') {
          localStorage.setItem('token', uuid);
          localStorage.setItem('userid', 123423);
          localStorage.setItem('username', 'Test1234');
          localStorage.setItem('feeling', 2);
          localStorage.setItem(
            'picture',
            'https://wallpapers.com/images/hd/cool-profile-picture-87h46gcobjl5e4xu.jpg'
          );
          this.loading = false;
          this.$router.push('/');
        }
      }
    }
  }
};
</script>

<template>
  <div class="login-page-background">
    <div class="login-page-container">
      <div class="login-page-title-container">
        <h1 class="login-page-title">WASAPhoto</h1>
        <div class="login-page-subtitle">
          <h3>Let's be</h3>
          <h3 class="login-page-colorful">colorful.</h3>
        </div>
      </div>
      <div class="login-page-login-form-container">
        <div class="login-page-login-form">
          <div class="login-page-login-form-title">Get started now.</div>
          <form class="login-page-login-form-input">
            <input type="text" placeholder="Username" v-model="username" />
            <button v-if="!loading" type="submit" @click="login">Login</button>
            <div v-if="loading" class="spinner">
              <div class="circle"></div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@keyframes rainbow {
  0% {
    background-position: 0% 50%;
  }
  100% {
    background-position: 100% 50%;
  }
}

.login-page-background {
  height: 100vh;
  width: 100vw;
  background: url('../assets/images/login-background.jpg');
  background-size: cover;
  .login-page-container {
    font-family: 'Courier New', Courier, monospace;
    display: flex;
    width: 100vw;
    height: 100vh;
    justify-content: center;
    align-items: center;
    flex-direction: row;
    backdrop-filter: blur(5px);
    @media screen and (max-width: 1250px) {
      flex-direction: column;
    }
    .login-page-title-container {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      width: 50%;
      height: 100vh;
      color: white;
      .login-page-title {
        font-size: 100px;
        margin-bottom: 0;
      }
      .login-page-subtitle {
        font-size: 40px;
        display: flex;
        flex-direction: row;
        .login-page-colorful {
          margin-left: 15px;
          background: linear-gradient(
            90deg,
            #ff6c3f,
            #ffa700,
            #ff0,
            #0f0,
            #00f,
            #a50,
            #ff00e0
          );
          background-size: 1000% 100%;
          -webkit-background-clip: text;
          color: transparent;
          animation: rainbow 8s linear infinite;
        }
      }
    }
    .login-page-login-form-container {
      width: 50%;
      height: 100vh;
      display: flex;
      justify-content: center;
      align-items: center;
      .login-page-login-form {
        border-radius: 20px;
        background: white;
        box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
        display: flex;
        justify-content: center;
        align-items: flex-start;
        flex-direction: column;
        padding: 40px 60px;
        .login-page-login-form-title {
          color: black;
          font-size: 22px;
          text-align: center;
        }
        .login-page-login-form-input {
          display: flex;
          flex-direction: row;
          align-items: center;
          input {
            width: 320px;
            height: 60px;
            border-radius: 14px;
            border: 1px solid #ccc;
            margin-top: 20px;
            padding: 0 15px;
            font-size: 1.2rem;
            font-family: 'Courier New', Courier, monospace;
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.08);
            &:focus {
              outline: none;
              border: 1px solid #ff6c3f;
            }
          }
          button {
            display: flex;
            justify-content: center;
            align-items: center;
            margin-top: 18px;
            margin-left: 20px;
            height: 58px;
            width: 50px;
            font-family: 'Courier New', Courier, monospace;
            border: none;
            background-color: #ffb074;
            border-radius: 10px;
            padding: 0px 40px;
            color: white;
            font-size: 18px;
            font-weight: bold;
            transition: all 0.3s ease;
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.08);
            &:hover {
              cursor: pointer;
              background-color: #ff882d;
              transform: scale(1.02);
            }
          }
        }

        .circle {
          width: 30px;
          height: 30px;
          border: 4px solid #ffa560; /* Color of the spinner */
          border-top: 4px solid transparent;
          border-radius: 50%;
          margin-top: 18px;
          margin-left: 20px;
          animation: spin 1.5s linear infinite;
        }

        @keyframes spin {
          0% {
            transform: rotate(0deg);
          }
          100% {
            transform: rotate(360deg);
          }
        }
      }
    }
  }
}
</style>
