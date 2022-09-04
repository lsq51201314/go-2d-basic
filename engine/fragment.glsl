#version 330 core
uniform sampler2D TEX_SAMPLER;

in vec2 fTexCoords;

const float lineWidth = 0.005;
const vec3  lineColor = vec3(1.0, 1.0, 1.0);
const float transparen = 1.0;

void main() 
{
    vec4 texColor = texture2D(TEX_SAMPLER,fTexCoords);
    // //透明
    // if (texColor.a != 0.0){
    //     texColor.a = transparen; 
    // } 
    //描边
    if (texColor.a == 0.0){
        float a = 
            texture2D(TEX_SAMPLER, vec2(fTexCoords.x + lineWidth, fTexCoords.y)).a +
			texture2D(TEX_SAMPLER, vec2(fTexCoords.x, fTexCoords.y - lineWidth)).a +
			texture2D(TEX_SAMPLER, vec2(fTexCoords.x - lineWidth, fTexCoords.y)).a +
			texture2D(TEX_SAMPLER, vec2(fTexCoords.x, fTexCoords.y + lineWidth)).a;

		if (texColor.a < 1.0 && a > 0.0)
			texColor = vec4(lineColor,transparen);
        else 
            discard;
    }
    // //圆形
    // vec2 center = vec2(0.5,0.5);
    // float radius = 0.4;
    // float res = sqrt(pow((fTexCoords.x - center.x),2.0)+pow((fTexCoords.y - center.y),2.0));  
    // if(res > radius) {
    //     discard;    
    // }  
    gl_FragColor = texColor;
}    