import { Paper } from '@material-ui/core';
import React from 'react';
import { connect } from 'react-redux';
import NetWorth from '../components/NetWorth';
import RecordTable from '../components/RecordTable';


class RecordTableContainer extends React.Component {
    componentDidMount() {
        this.props.fetchRecords()
    }

    render() {
        return (
            <Paper style={{ width: "75%", margin: 'auto' }}>
                <NetWorth isLoading={this.props.isLoading} netWorth={this.props.records.net_worth} />
                <RecordTable
                    records={this.props.records}
                    isLoading={this.props.isLoading}
                    error={this.props.error}
                    createRecord={this.props.createRecord}
                    deleteRecord={this.props.deleteRecord}
                />
            </Paper>
        )
    }
}

const mapStateToProps = state => ({
    records: state.record.records,
    isLoading: state.record.isLoading,
    error: state.record.error,
})

const mapDispatchToProps = ({ record: { fetchRecords, createRecord, deleteRecord } }) => ({
    fetchRecords: () => fetchRecords(),
    createRecord: ({ name, type, balance }) => createRecord({ name, type, balance }),
    deleteRecord: ({ id }) => deleteRecord({ id }),
})

export default connect(mapStateToProps, mapDispatchToProps)(RecordTableContainer);

