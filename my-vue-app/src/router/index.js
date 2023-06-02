import Question from "../components/Question.vue";
import QuestionDetails from "../components/QuestionDetails.vue";
import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import MainPage from "../components/MainPage.vue";

const routes = [
    // 动态字段以冒号开始
    { path: '/questionDetails/:id',name:"question.details", component: QuestionDetails, props: true },
    {path: '/',component:  MainPage}
]


const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHashHistory(),
    routes, // `routes: routes` 的缩写
})

export default  router