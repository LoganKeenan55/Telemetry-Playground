function showError(errorText){
    const errorBoxDiv = document.getElementById('error-box');
    const errorTextElement = document.createElement('p');
    errorTextElement.innerText = errorText;
    errorBoxDiv.appendChild(errorTextElement);
    console.log(errorText);
}


function helloTriangle(){
    /**@type {HTMLCanvasElement|null} */
    const canvas = document.getElementById('demo-canvas');
    if(!canvas){
        showError("Can't find demo canvas reference :(");
        return;
    }
    const gl = canvas.getContext('webgl2'); 
    if(!gl){
        showError("Browser doesn't support webGL2 :(");
        return;
    }

    const triangleVertices = [
        0.0, 0.5,   1.0,0.0,0.0,
        -0.5, -0.5, 0.0,1.0,0.0,
        0.5, -0.5,  0.0,0.0,1.0
    ];
    
    const verticesCPUBuffer = new Float32Array(triangleVertices);

    const vertexBuffer = gl.createBuffer();
    gl.bindBuffer(gl.ARRAY_BUFFER,vertexBuffer);
    gl.bufferData(gl.ARRAY_BUFFER,verticesCPUBuffer,gl.STATIC_DRAW);

    const vertexShaderCode = `#version 300 es
    precision mediump float;
    
    layout(location = 0) in vec2 aPos;
    layout(location = 1) in vec3 aColor;

    out vec3 color;

    void main(){
        color = aColor;
        gl_Position = vec4(aPos, 0.0, 1.0);

    }`;


    const vertexShader = gl.createShader(gl.VERTEX_SHADER);
    gl.shaderSource(vertexShader,vertexShaderCode);
    gl.compileShader(vertexShader);
    if(!gl.getShaderParameter(vertexShader, gl.COMPILE_STATUS)){
        const compileError = gl.getShaderInfoLog(vertexShader);
        showError(`Failed to Compile vertex shader - ${compileError}`);
    }


    const fragmentShaderCode = `#version 300 es
    precision mediump float;
    
    in vec3 color;

    out vec4 fragColor;

    void main(){
        fragColor = vec4(color,1.0);
    
    }`; 
    const fragmentShader = gl.createShader(gl.FRAGMENT_SHADER);
    gl.shaderSource(fragmentShader,fragmentShaderCode);
    gl.compileShader(fragmentShader);
    if(!gl.getShaderParameter(fragmentShader, gl.COMPILE_STATUS)){
        const compileError = gl.getShaderInfoLog(fragmentShader);
        showError(`Failed to Compile fragement shader - ${compileError}`);
        return;
    }

    const program = gl.createProgram();
    gl.attachShader(program,vertexShader);
    gl.attachShader(program,fragmentShader);

    gl.linkProgram(program);
    
    canvas.width = canvas.clientWidth;
    canvas.height = canvas.clientHeight;

    gl.clearColor(0.08,0.08,0.08,1.0);
    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

    //rasterizer 
    gl.viewport(0,0,canvas.width,canvas.height);

    //set up GPU program
    gl.useProgram(program);
    gl.enableVertexAttribArray(0);
    gl.enableVertexAttribArray(1);

    //input assembler
    gl.bindBuffer(gl.ARRAY_BUFFER,vertexBuffer);
    gl.vertexAttribPointer(0,2,gl.FLOAT,false,5*Float32Array.BYTES_PER_ELEMENT,0)
    gl.vertexAttribPointer(1,3,gl.FLOAT,false,5*Float32Array.BYTES_PER_ELEMENT,2*Float32Array.BYTES_PER_ELEMENT)

    gl.drawArrays(gl.TRIANGLES,0,3);
}




try{
    helloTriangle();
}catch(e){
    showError(`uncaught javascript exception: ${e}`);
}