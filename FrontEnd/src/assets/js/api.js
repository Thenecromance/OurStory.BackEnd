import axios from 'axios';


export async function initResource(store ) {
  try {
   await axios.get('/agronDash/title')
   .then(response =>{
    console.log(response.data);
    store.commit('setPageTitle', response.data.title);
   });
   
  } catch (error) {
    // console.error(error);
  }
}