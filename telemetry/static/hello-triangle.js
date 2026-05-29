function showError(errorText){
    const errorBoxDiv = document.getElementById('error-box');
    const errorTextElement = document.createElement('p');
    errorTextElement.innerText = errorText;
    errorBoxDiv.appendChild(errorTextElement);
    console.log(errorText);
}

showError("you stupid!");

function helloTriangle(){
    const canvas = document.getElementById('demo-canvas');
    if(!canvas){
        showError("Can't find demo canvas reference :(");
    }
    
}



try{
    helloTriangle();
}catch(e){
    showError(`uncaught javascript exception: ${e}`);
}