/**
* @auther: kid1999
* @date: 2021/7/14 19:30
* @desciption:  Login
*/
<template>
  <div class="q-pa-md">
    <div class="row">
      <div class="offset-xs-3 col-xs-10 col-sm-8 col-md-6">
        <q-card class="my-card">
          <q-card-section>
            <div class="text-h5">登录</div>
          </q-card-section>
          <q-card-section>
            <q-form
              @submit="onSubmit"
              @reset="onReset"
              class="q-gutter-md"
            >
              <q-input
                filled
                v-model="StudentID"
                label="学号"
                hint="student ID"
                lazy-rules
                :rules="[ val => val && val.length > 0 || '请输入学号']"
              />

              <q-input
                filled
                v-model="Password"
                label="密码"
                hint="password"
                lazy-rules
                :rules="[ val => val && val.length > 0 || '请输入密码']"
              />

              <q-btn label="登 录" type="submit" color="primary"/>
              <q-btn label="重新填写" type="reset" color="primary" flat class="q-ml-sm" />

            </q-form>
          </q-card-section>
        </q-card>
      </div>
    </div>

  </div>
</template>

<script>
import { post } from '../util/request'
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'Login',
  data () {
    return {
      StudentID: '',
      Password: ''
    }
  },
  methods: {
    onReset () {
      this.StudentID = ''
      this.Password = ''
    },
    onSubmit () {
      if (this.StudentID === '' || this.Password === '') {
        this.$q.notify({
          type: 'warning',
          message: '请输入学号密码！'
        })
        return
      }
      post('/login', { StudentID: this.StudentID, Password: this.Password })
        .then(res => {
          console.info(res)
          localStorage.setItem('token', res.token)
          localStorage.setItem('user', res.user)
          this.$q.notify({
            type: 'success',
            message: '登录成功！'
          })
        })
    }
  }
})

</script>

<style scoped>

</style>
