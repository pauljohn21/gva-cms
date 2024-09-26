
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="company字段:" prop="company">
          <el-input v-model="formData.company" :clearable="true"  placeholder="请输入company字段" />
       </el-form-item>
        <el-form-item label="统一社会信用代码:" prop="code">
          <el-input v-model="formData.code" :clearable="true"  placeholder="请输入统一社会信用代码" />
       </el-form-item>
        <el-form-item label="开户行:" prop="bank">
          <el-input v-model="formData.bank" :clearable="true"  placeholder="请输入开户行" />
       </el-form-item>
        <el-form-item label="对公账户:" prop="account">
          <el-input v-model="formData.account" :clearable="true"  placeholder="请输入对公账户" />
       </el-form-item>
        <el-form-item label="地址:" prop="addrs">
          <el-input v-model="formData.addrs" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="电话:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入电话" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入状态" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="signature字段:" prop="signature">
          <el-input v-model="formData.signature" :clearable="true"  placeholder="请输入signature字段" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPolicyholder,
  updatePolicyholder,
  findPolicyholder
} from '@/api/cms/policyholder'

defineOptions({
    name: 'PolicyholderForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            company: '',
            code: '',
            bank: '',
            account: '',
            addrs: '',
            phone: '',
            status: '',
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
            signature: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPolicyholder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPolicyholder(formData.value)
               break
             case 'update':
               res = await updatePolicyholder(formData.value)
               break
             default:
               res = await createPolicyholder(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
