import express from "express";

const app = express()
const port = 8000

app.use(express.json());
app.use(express.urlencoded({extended: true}));


app.get('/', (req, res)=>{
    return res.status(200).send("Welcome to My server");
});

app.post('/post', (req, res) => {
    const myJson = req.body;
    res.status(200).send(myJson);
});

app.post('/postForm', (req, res) => {
    return res.status(200).send(JSON.stringify(req.body));
})

app.listen(port, () => {
    console.log(`Server runnning at http://localhost:${port}`);
})