// constante consumption
const low = 100;
const hight = 300;

function color(value){
    console.log(value)
    if(value < low){
    // value is low
        return 'green';
    }else if((value >= low) && (value <= hight)){
    // value is in the middle
        return 'white';
    }else{
    // value is hight
        return 'red';
    }
}

