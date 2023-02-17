const fs = require('fs');

loadData();

async function loadData() {
    let data = await fs.promises.readFile("test.txt");
    let text = data.toString();
    let i = 10;
    data.toString().split('\n')[0].split(" ").map(e => {
        return {
            n: e,
            w: isNaN(e)
        }
    }).every(e => {
        if (e.w) {
            text = text.replaceAll(e.n, i++);
        }        
        return true
    })
    text = text.replaceAll("\n", "},\n{").replaceAll(" ", ",");
    console.log("{" + text + "}");
}
