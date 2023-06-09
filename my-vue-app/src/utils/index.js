import {useStore} from "vuex";

export function handleTime(time){
    return time.replace("T"," ").split("+")[0]
}

export function handleComments(comments){
    while (already(comments)) {
        comments.forEach((item)=> {
            if (item.BeID !== item.QuestionID && !item.already ){
                comments.forEach(titem=> {
                    if (titem.ID === item.BeID){
                        if (titem.Be == undefined){
                            titem.Be =[]
                        }
                        titem.Be.push(item)
                        item["already"] = true
                    }
                })
            }
        })

    }

    let newcomments = comments.filter((item)=>{
        return !item.already
    })

    return newcomments

}

function already(comments){
    for (let commentsKey of comments) {
        if (commentsKey.BeID !== commentsKey.QuestionID && !commentsKey.already){
            return true
        }
    }

    return false
}

export function handleResult(result){
    if (result.data.status == 500){
        ElMessage({
            message:result.data.messae,
            type:"warning"
        })

        return
    }

    if(result.data.status == 200){
        ElMessage({
            message:result.data.message,
            type:"success"
        })
    }

    setTimeout(()=>{
        location.reload()
    },500)
}


export function getQuestion(ID){
    const store = useStore()
    for (let i = 0; i < store.state.allQuestions.length; i++) {
        if (store.state.allQuestions[i].ID == ID ){
            return store.state.allQuestions[i].Title
        }
    }

}

export function getNumberOfComments(comments) {
    let sum = 0
    for (const commentsKey in comments) {
        sum += comments[commentsKey].length
    }
    return sum
}