export function handleTime(time){
    return time.replace("T"," ").split("+")[0]
}

export function handleComments(comments){
    while (already(comments)) {
        comments.forEach((item)=> {
            if (item.BeID !== item.QuestionID){
                comments.forEach(titem=> {
                    if (titem.ID === item.BeID){
                        if (titem.Be == undefined){
                            titem.Be =[]
                        }
                        titem.Be.push(item)
                        let index = comments.indexOf(item)
                        comments.splice(index,1)
                    }
                })
            }
        })

    }

}

function already(comments){
    for (let commentsKey of comments) {
        if (commentsKey.BeID !== commentsKey.QuestionID){
            return true
        }
    }

    return false
}