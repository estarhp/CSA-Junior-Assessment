export function handleTime(time){
    return time.replace("T"," ").split("+")[0]
}