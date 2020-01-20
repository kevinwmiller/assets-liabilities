import axios from 'axios';

const record = {
    state: {
        isLoading: false,
        error: null,
        records: []
    }, // initial state
    reducers: {
        setIsLoading(state, payload) {
            return {
                ...state,
                isLoading: payload,
                error: null,
            }
        },
        setError(state, payload) {
            return {
                ...state,
                isLoading: false,
                error: payload,
            }
        },
        setRecords(state, payload) {
            return {
                ...state,
                isLoading: false,
                error: null,
                records: payload,
            }
        },
    },
    effects: dispatch => ({
        // handle state changes with impure functions.
        // use async/await for async actions
        fetchRecords: async (payload, rootState) => {
            console.log("Fetching records")
            let res = await axios.get("/finances/records");
            console.log(res)
            dispatch.record.setRecords(res.data);
        },
        createRecord: async (payload, rootState) => {
        },
        updateRecord: async (payload, rootState) => {
        },
        deleteRecord: async (payload, rootState) => {
        },
    })
}

export default record
