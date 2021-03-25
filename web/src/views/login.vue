<template>
  <div class="loginStyle">
    <a-form :model="loginForm" class="loginContainer">
      <h3 class="loginTitle">系统登录</h3>
      <a-form-item prop="username">
        <a-input
            size="normal"
            type="text"
            v-model="loginForm.username"
            auto-complete="off"
            placeholder="请输入用户名"
        ></a-input>
      </a-form-item>
      <a-form-item prop="password">
        <a-input
            size="normal"
            type="password"
            v-model="loginForm.password"
            auto-complete="off"
            placeholder="请输入密码"
        ></a-input>
      </a-form-item>
      <a-button
          size="normal"
          type="primary"
          style="width: 100%;"
          @click="submitLogin">登录
      </a-button
      >
    </a-form>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      loading: false,
      loginForm: {
        username: 'gj',
        password: '123',
      },
      checked: true,
      rules: {
        username: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
        ],
        password: [{required: true, message: '请输入密码', trigger: 'blur'}],
      },
    }
  },
  methods: {
    submitLogin() {
      let path = this.$route.query.redirect
      this.$router.replace(path == '/' || path == undefined ? '/home' : path)
    },
  },
}
</script>

<style>
.loginStyle {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  width: 100%;
  height: 100%;
  background-image: url('../assets/bg.jpg');
  background-position: 100%;
  /*background-size:auto auto;*/
  background-size: cover;
}

.loginContainer {
  border-radius: 15px;
  background-clip: padding-box;
  margin: 180px auto;
  width: 350px;
  padding: 15px 35px 15px 35px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}

.loginTitle {
  margin: 15px auto 20px auto;
  text-align: center;
  color: #505458;
}

.loginRemember {
  text-align: left;
  margin: 0px 0px 15px 0px;
}
</style>