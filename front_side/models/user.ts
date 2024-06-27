
export enum UserRole {
    USER,
    ADMIN,
    MASTER
}

export type User = {
    id: number;         // user id 
    username: string;   // user name 
    avatar?: string;     // user avatar url. if not set it will be a default avatar
    nickname: string;   // user nickname which will  display to others
    role: number;       // user role 0:user 1:admin 2:master see above 
    email: string;      // user email
    mbti?: string;       // mbti type , not necessary 
    birthday?: number;   // unix time stamp
    gender?: string;     // as what he want to be 
    created_at: number; // unix time stamp 
    last_login: number; // unix time stamp
};