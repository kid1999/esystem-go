/**
* @auther: kid1999
* @date: 2021/7/14 20:29
* @desciption:  Register
*/
<template>
  <div class="q-pa-md">
    <div class="row">
      <div class="offset-xs-3 col-xs-10 col-sm-8 col-md-6">
        <q-card class="my-card">
          <q-card-section>
            <div class="text-h5">注册</div>
          </q-card-section>
          <q-card-section>
            <q-form
              @submit="onSubmit"
              @reset="onReset"
              class="q-gutter-md"
              ref="myForm"
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
                v-model="Username"
                label="姓名"
                hint="user name"
                lazy-rules
                :rules="[ val => val && val.length > 0 || '请输入姓名']"
              />

              <q-input
                filled
                v-model="Password"
                label="密码"
                hint="password"
                lazy-rules
                :rules="[ val => val && val.length > 0 || '请输入密码']"
              />

              <q-input
                filled
                v-model="Password2"
                type="passoword"
                label="确认密码"
                hint="password"
                lazy-rules
                :rules="[ val => val && val === Password || '前后密码不一致']"
              />

              <q-input
                filled
                v-model="Phone"
                label="手机"
                mask="### #### ####"
                unmasked-value
                hint="请输入正确的手机号"
                lazy-rules
                :rules="[ val => val && val.length === 11 || '请输入正确的手机号']"
              />

              <q-input
                filled
                v-model="Wx"
                label="微信"
                hint="wechat ID"
              />

              <q-input
                filled
                v-model="QQ"
                label="QQ"
                hint="QQ"
              />

              <q-input
                filled
                v-model="Email"
                label="邮箱"
                hint="email"
              />

              <q-btn label="提 交" type="submit" color="primary" />
              <q-btn label="重新填写" type="reset" color="primary" flat class="q-ml-sm"/>

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
  name: 'Register',
  data () {
    return {
      StudentID: '',
      Username: '',
      Password: '',
      Password2: '',
      Phone: '',
      Wx: '',
      QQ: '',
      Email: ''
    }
  },
  methods: {
    onReset () {
      this.StudentID = ''
      this.Username = ''
      this.Password = ''
      this.Password2 = ''
      this.Phone = ''
      this.Wx = ''
      this.QQ = ''
      this.Email = ''
    },
    onSubmit () {
      this.$refs.myForm.validate().then(success => {
        if (success) {
          post('/register', { StudentID: this.StudentID, Password: this.Password, Username: this.Username, Phone: this.Phone, Wx: this.Wx, QQ: this.QQ, Email: this.Email })
            .then(res => {
              console.info(res)
            })
        } else {
          this.$q.notify({
            type: 'warning',
            message: '请填写必要信息！'
          })
        }
      })
    }
  }
})
</script>

<style scoped>

</style>
