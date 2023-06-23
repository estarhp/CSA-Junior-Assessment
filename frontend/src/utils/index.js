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

export function handleResult(result,willreload=true){
    if (result.data.status == 500){
        ElMessage({
            message:result.data.message,
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
if (willreload) {
    setTimeout(()=>{
        location.reload()
    },500)
}

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

export function reload(result){
    if (result.data.status === 200){
        setTimeout(()=>{
            location.reload()
        },1000)
    }
}

const color = ["#F0FFFF", "#FAFAD2", "#F5FFFA", "#F0FFF0", "#FFFAF0", "#FFFFF0", "#FFFACD", "#F8F8FF", "#F5F5F5", "#F0F8FF", "#F0E68C", "#EEE8AA", "#E0FFFF", "#E6E6FA", "#DCDCDC", "#D3D3D3", "#D8BFD8", "#DDA0DD", "#DC143C", "#DA70D6", "#DB7093", "#DAA520", "#D2691E", "#CD853F", "#C71585", "#BDB76B", "#BC8F8F", "#BA55D3", "#B8860B", "#AFEEEE"]

export function getRandomColor() {
    const randomIndex = Math.floor(Math.random() * color.length);
    return color[randomIndex];
}