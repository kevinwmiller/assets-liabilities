import axios from 'axios';

const record = {
    state: {
        isLoading: false,
        error: null,
        records: {}
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
            dispatch.record.setIsLoading(true);
            try {
                let res = await axios.get("/finances/records");
                console.log(res)
                dispatch.record.setRecords(res.data);
            } catch (e) {
                console.log(e)
                dispatch.record.setError(e);
            }
        },
        createRecord: async (payload, rootState) => {
            try {
                let res = await axios.post("/finances/records", payload);
                dispatch.record.fetchRecords()
            } catch (e) {
                console.log(e)
                dispatch.record.setError(e);
            }
        },
        deleteRecord: async (payload, rootState) => {
            try {
                await axios.delete(`/finances/records/${payload.id}`);
                dispatch.record.fetchRecords()
            } catch (e) {
                console.log(e)
                dispatch.record.setError(e);
            }
        },
    })
}

export default record
