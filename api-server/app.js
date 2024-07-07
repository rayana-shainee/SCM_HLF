const express=require("express");
const app= express();
const bodyParser=require("body-parser");
const { registerUser } = require("./registerUser");
const { tx } = require("./tx");
const { query } = require("./query");
const cors = require('cors');

app.use(cors()); // Enable CORS
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.static('public'));
app.listen(4000,()=>{
    console.log("server started");

})

app.post("/register",async(req,res)=>{

    try {
        let orgMSP=req.body.orgMSP;
        let userId=req.body.userId;
        let result =await registerUser({OrgMSP:orgMSP,userId:userId});
        res.send(result);

    } catch (error) {
        res.send(error) 
    }
});


app.post("/tx",async(req,res)=>{
    try {
        var request={
            chaincodeName:req.body.chaincodeName,
            channelName:req.body.channelName,
            userId:req.body.userId,
            org:req.body.orgMSP,
            data:req.body.data
        };
       let result= await tx(request);
       res.send(result)
        
    } catch (error) {
        res.send(error);
    }

})

app.post("/query", async (req, res) => {
    try {
        var request = {
            chaincodeName: req.body.chaincodeName,
            channelName: req.body.channelName,
            userId: req.body.userId,
            org: req.body.orgMSP,
            data: req.body.data
        }
        let result = await query(request);
        res.send(result)
        
    } catch (error) {
        res.send(error);
    }
});
