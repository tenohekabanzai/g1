const express = require('express')
const app = express()

app.use(express.json())
app.use(express.urlencoded({extended:true}))

app.get('/',(req,res)=>{
    res.status(200).send("Welcome to server")
})

app.get('/get',(req,res)=>{
    res.status(200).json({message:"Hello from abc.com"})
})

app.post('/post',(req,res)=>{
    let myJson = req.body
    console.log(myJson)
    res.status(200).send(myJson)
})

app.post('/postform',(req,res)=>{
    res.status(200).send(JSON.stringify(req.body))
})

app.listen(5004,(req,res)=>{
    console.log("App running on PORT 5004")
})


