import {ref} from 'vue';

const getAllPosts = () => {
    const oPosts = ref(null);
    const oError = ref(null);

    const fnLoad = async () => {
        try{
            const oData = await fetch('http://jsonplaceholder.typicode.com/posts');
            if(!oData.ok) {
                throw Error('No data avbailable');
            }   
            oPosts.value = await oData.json();
        } catch(oErr) {
            oError.value = oErr.message;
            console.error(oErr.value);
        }
    }
    return {oPosts, oError, fnLoad};
};

export default getAllPosts;