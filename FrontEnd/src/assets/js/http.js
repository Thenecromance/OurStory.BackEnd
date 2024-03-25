import axios from "axios";

function get(url) {

    const response = axios.get(url)
    return response;
}