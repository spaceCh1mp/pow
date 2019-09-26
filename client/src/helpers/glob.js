export var Glob = {
    URL : getPath()
}

function getPath(){
    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
        return 'http://localhost:3000/'
    }
    return 'http://pow.kenechukwuagugua.com/'
}