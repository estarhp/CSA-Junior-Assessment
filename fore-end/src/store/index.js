import {createStore} from "vuex";
import axios from "axios";
import {ElMessage} from "element-plus";
import {reactive, ref, toRefs} from "vue";


const state = reactive({
    isLogin:"",
    userDetails:"",
    allQuestions:"",
    initData:{
        telephone:"",
        address:""
    }
})

const mutations ={
  isLogin(state,message){
      state.isLogin = message

      if (message == "false") {
          ElMessage({
              message:"现在是游客状态",
              type:"warning"
          })
      }
  },
    userDetails(state,details){
      details.Follows = JSON.parse(details.Follows)
        console.log( details)
      state.userDetails = details
      state.initData.telephone = details.Telephone
      state.initData.address = details.Address
        if (location.hash !== "#/"){
            return
        }
        ElMessage({
            message:"你好! "+details.Username,
            type:"success"
        })

    },
    getAllQuestions(state,questions){
      state.allQuestions = ref(questions)
    }
}

const actions = {
    AlreadyLogin: async function (context) {
        const data = await axios({
            url: "/api/isLogin"
        }).catch(err => {
            console.error(err)
            return
        })

        context.commit("isLogin", data.data.message)
        if (data.data.message === "true"){
            await context.dispatch("getUserDetails")
        }

    },
    getUserDetails: async function (context) {
        const result = await axios({
            url: "/api/userDetails",
            method: "get"
        })

        context.commit("userDetails", result.data.userDetails)
    },
    async addComment(context, object) {

        if (object.content.value == "") {
            ElMessage({
                message: "输入不合法",
                type: "warning"
            })
            return
        }

        const formData = new FormData()
        formData.append("content", object.content)
        formData.append("beID", object.beID)
        formData.append("questionID", object.questionID)
        formData.append("beUsername", object.beUsername || "")
        const result = await axios({
            url: "/api/addComment",
            method: "POST",
            data: formData
        }).catch(err => {
            ElMessage({
                message: result.data.message,
                type: 'warning'
            })
            return false
        })

        if (result.data.status == 500) {
            ElMessage({
                message: result.data.message,
                type: "warning"
            })
            return false
        }

        ElMessage({
            message: result.data.message
        })
        return true

    },
    async getAllQuestions(context) {

        const result = await axios({
            url: "/api/getAllQuestions",
            method: "get",
        })

        context.commit("getAllQuestions", result.data.questions)

    }
}

export default createStore({
    state,
    actions,
    mutations
})