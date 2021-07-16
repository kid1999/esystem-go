/**
* @auther: kid1999
* @date: 2021/7/15 10:51
* @desciption:  AddGoods
*/
<template>
  <div class="q-pa-md">
    <div class="row">
      <div class="offset-xs-3 col-xs-10 col-sm-8 col-md-6">
        <q-card class="my-card">
          <q-card-section>
            <div class="text-h5">上架商品</div>
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
                v-model="GoodsName"
                label="商品名"
                hint="Goods Name"
                lazy-rules
                :rules="[ val => val && val.length <= 30 && val.length > 0 || '请输入商品名(最多30个字符)']"
              />

              <q-input
                filled
                autogrow
                v-model="Description"
                label="商品描述"
                hint="Description"
                lazy-rules
                :rules="[ val => val && val.length <= 200 && val.length > 0 || '请输入商品描述(最多200个字符)']"
              />

              <q-input
                filled
                v-model="Username"
                label="所有者"
                hint="user name"
                readonly
              />

              <q-input
                filled
                v-model="Price"
                prefix="¥"
                label="预估价值"
                hint="predict price"
                lazy-rules
                :rules="[ val => val && val <= 1000000 && val >= 0 || '请输入商品预估价格(最多100W)']"
              />

              <q-input
                filled
                v-model="Phone"
                label="手机"
                type="tel"
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
                type="email"
                hint="email"
              />

              <q-uploader
                :factory="factoryFn"
                color="teal"
                flat
                max-files="1"
                extensions=".gif,.jpg,.jpeg,.png"
                style="max-width: 300px"
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
import { successMsg, failMsg } from '../util/msg'
export default defineComponent({
  name: 'AddGoods',
  data () {
    return {
      Price: 0.0,
      GoodsName: '',
      Description: '',
      OwnerID: '',
      UserName: '',
      Phone: '',
      Wx: '',
      QQ: '',
      ImgUrl: ''
    }
  },
  created () {
    const localUser = JSON.parse(localStorage.getItem('user'))
    console.info(localUser)
    if (localUser == null) {
      failMsg('请先登录！')
      this.router.push('/login')
    }
    this.Username = localUser.Username
    this.OwnerStudentID = localUser.StudentID
    this.Wx = localUser.Wx
    this.Phone = localUser.Phone
    this.QQ = localUser.QQ
    this.Email = localUser.Email
  },
  computed: {
    ComputePirce () {
      const num = parseFloat(this.Price)
      if (num.toString() === 'NaN') {
        return 0.0
      } else {
        return num
      }
    }
  },
  methods: {
    factoryFn (files) {
      const fd = new FormData()
      fd.append('file', files[0])
      const _this = this
      post('/upload', fd, {
        headers: { 'Content-Type': undefined }
      }).then(function (res) {
        _this.ImgUrl = res
        console.info(res)
        successMsg('图片上传成功')
      })
    },
    onReset () {
      this.GoodsName = ''
      this.Price = ''
      this.Phone = ''
      this.Wx = ''
      this.QQ = ''
      this.Email = ''
      this.Description = ''
    },
    onSubmit () {
      this.$refs.myForm.validate().then(success => {
        if (success) {
          post('/goods', { ImgUrl: this.ImgUrl, OwnerStudentID: this.OwnerStudentID, GoodsName: this.GoodsName, Username: this.Username, Phone: this.Phone, Wx: this.Wx, QQ: this.QQ, Email: this.Email, Description: this.Description, Price: this.ComputePirce })
            .then(res => {
              console.info(res)
              successMsg('上架成功！')
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
