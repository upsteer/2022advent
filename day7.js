var tree = {
    "/": {
        "dir.a":{
            "dir.e": {
                "i": 584
            },
            "f": 29116,
            "g": 2557,
            "h.lst": 62596
        },
        "b.txt": 14848514,
        "c.dat": 8504156,
        "dir.d":{
            "j": 4060174,
            "d.log": 8033020,
            "d.ext": 5626152,
            "k": 7214296
        }
    }
}

var summs = []
function traverse(tree){
    var keys = Object.keys(tree);
    var sum = 0
    for (var i=0; i<keys.length;i++){
        // console.log(keys[i], tree[keys[i]])
        subtree = tree[keys[i]]
        if(typeof subtree === 'object' && subtree !== null){
            console.log(keys[i], subtree)
            sum+=traverse(subtree)
        } else {
            sum += subtree
        }
    }
    summs.push(sum)
    return sum
}

traverse(tree["/"])
console.log(summs)
