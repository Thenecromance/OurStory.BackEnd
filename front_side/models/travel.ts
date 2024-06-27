
export enum TravelState {
    PENDING,
    ONGOING,
    COMPLETED
}


// travel model
export type Travel = {
    id: number;         //travel id 
    state: number;      //travel state 0:pending 1:ongoning 2:completed 
    owner: number;      //who created this travel 
    start: number;      //when the travel start 
    end: number;        //when the travel end 
    location: string;   //where do you go 
    details: string;    //travel details just like a memo. maybe memo can be a better name
    together: [];       //who will go with you (if you take this travel alone, this field will be empty)
    img: string;        //the most beautiful photos todo: the backend I haven't implemented the image upload function yet
};