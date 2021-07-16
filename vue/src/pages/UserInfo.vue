/**
* @auther: kid1999
* @date: 2021/7/15 8:10
* @desciption:  UserInfo
*/
<template>
  <div class="q-pa-md">
    <div class="row">
      <div class="offset-xs-3 col-xs-10 col-sm-8 col-md-6">
        <q-card class="my-card">
          <q-card-section>
            <div class="text-h5">个人信息</div>
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
                v-model="user.StudentID"
                label="学号"
                hint="student ID"
                readonly
              />

              <q-input
                filled
                v-model="user.Username"
                label="姓名"
                hint="user name"
                readonly
              />

              <q-input
                filled
                v-model="user.Phone"
                label="手机"
                clearable
                type="tel"
                mask="### #### ####"
                unmasked-value
                hint="请输入正确的手机号"
                lazy-rules
                :rules="[ val => val && val.length === 11 || '请输入正确的手机号']"
              />

              <q-input
                filled
                v-model="user.Wx"
                label="微信"
                hint="wechat ID"
                clearable
              />

              <q-input
                filled
                v-model="user.QQ"
                label="QQ"
                hint="QQ"
                clearable
              />

              <q-input
                filled
                v-model="user.Email"
                label="邮箱"
                hint="email"
                type="email"
                clearable
              />

              <q-btn label="保存修改" type="submit" color="primary" />
              <q-btn label="重新填写" type="reset" color="primary" flat class="q-ml-sm"/>
              <q-btn label="修改密码" color="primary" @click="prompt = true" />

              <q-dialog v-model="prompt" persistent>
                <q-card style="min-width: 350px">
                  <q-card-section>
                    <div class="text-h6">修改密码</div>
                  </q-card-section>

                  <q-card-section class="q-pt-none">

                    <q-input v-model="user.Password" label="旧密码" filled :type="isPwd1 ? 'password' : 'text'" hint="password" lazy-rules :rules="[ val => val && val.length > 0 || '请输入密码']">
                      <template v-slot:append>
                        <q-icon
                          :name="isPwd1 ? 'visibility_off' : 'visibility'"
                          class="cursor-pointer"
                          @click="isPwd1 = !isPwd1"
                        />
                      </template>
                    </q-input>

                    <q-input v-model="Password" label="新密码" filled :type="isPwd2 ? 'password' : 'text'" hint="password" lazy-rules :rules="[ val => val && val.length > 0 || '请输入密码']">
                      <template v-slot:append>
                        <q-icon
                          :name="isPwd2 ? 'visibility_off' : 'visibility'"
                          class="cursor-pointer"
                          @click="isPwd2 = !isPwd2"
                        />
                      </template>
                    </q-input>

                    <q-input v-model="Password2" label="确认密码" filled :type="isPwd3 ? 'password' : 'text'" hint="password" lazy-rules :rules="[ val => val && val === Password || '前后密码不一致']">
                      <template v-slot:append>
                        <q-icon
                          :name="isPwd3 ? 'visibility_off' : 'visibility'"
                          class="cursor-pointer"
                          @click="isPwd3 = !isPwd3"
                        />
                      </template>
                    </q-input>

                  </q-card-section>

                  <q-card-actions align="right" class="text-primary">
                    <q-btn flat label="取消" v-close-popup />
                    <q-btn flat label="保存修改" v-close-popup />
                  </q-card-actions>
                </q-card>
              </q-dialog>

            </q-form>
          </q-card-section>
        </q-card>
      </div>
    </div>

  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { put } from '../util/request'
import { successMsg, failMsg } from '../util/msg'
export default defineComponent({
  name: 'UserInfo',
  data () {
    return {
      user: {},
      newPassword: 'change pass',
      prompt: false,
      Password: '',
      Password2: '',
      isPwd1: true,
      isPwd2: true,
      isPwd3: true
    }
  },
  created () {
    const localUser = JSON.parse(localStorage.getItem('user'))
    console.info(localUser)
    if (localUser == null) {
      failMsg('请先登录！')
      this.router.push('/login')
    }
    this.user = localUser
  },
  methods: {
    onReset () {
      const studentId = this.user.StudentID
      const username = this.user.Username
      this.user = {}
      this.user.StudentID = studentId
      this.user.Username = username
    },
    onSubmit () {
      this.$refs.myForm.validate().then(success => {
        if (success) {
          put('/user', this.user)
            .then(res => {
              console.info(res)
              localStorage.setItem('user', JSON.stringify(res))
              successMsg('修改成功！')
            })
        } else {
          failMsg('请填写必要信息！')
        }
      })
    }
  }
})
</script>

<style scoped>

</style>
