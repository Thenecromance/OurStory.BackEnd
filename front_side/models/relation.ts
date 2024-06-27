
export enum RelationType {
    UNKNOWN,
    FRIEND,     // friend link can link multiple people
    COUPLE,     // couple link only can link one person
}

export type Relation = {
    id: number;
    user_id: number;
    friend_id: number;
    relation_type: number;
    stamp: number;
}

export type RelationList = {
    [key: number]: Relation;
}

export enum RelationOperation {
    Bind,
    Unbind,
    Modify,
}

export type RelationHistory = {
    id: number;         // history id 
    user_id: number;    // which user do this operation
    operation_type: number; // 0: add 1: delete 2: update
    operation_time: number; // unix time stamp
    operation: number;  // which relation do this operation
    target_id: number;  // the target of this operation
};