import {createStore} from "vuex";
import axios from "axios";


const state = {
  isLogin:"",
  userDetails:"",
  allQuestions:""
}

const mutations ={
  isLogin(state,message){
      state.isLogin = message
  },
    userDetails(state,details){
      state.userDetails = details
    },
    getAllQuestions(state,questions){
      state.allQuestions = questions
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
        await context.dispatch("getUserDetails")

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