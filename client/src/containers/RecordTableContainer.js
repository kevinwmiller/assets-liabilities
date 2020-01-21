import { Typography } from '@material-ui/core';
import React from 'react';
import { connect } from 'react-redux';
import NetWorth from '../components/NetWorth';
import RecordTable from '../components/RecordTable';
class RecordTableContainer extends React.Component {
    componentDidMount() {
        this.props.fetchRecords()
    }

    render() {
        if (!this.props.records.asset_total) {
            this.props.records.asset_total = 0
        }
        if (!this.props.records.liability_total) {
            this.props.records.liability_total = 0
        }

        return (
            <div style={{ maxWidth: "1000px", margin: 'auto' }}>
                <NetWorth isLoading={this.props.isLoading} netWorth={this.props.records.net_worth} />
                <RecordTable
                    records={this.props.records}
                    isLoading={this.props.isLoading}
                    error={this.props.error}
                    createRecord={this.props.createRecord}
                    deleteRecord={this.props.deleteRecord}
                />
                <div>
                    <Typography style={{ marginRight: '20px' }} variant='subtitle1' component='span'>
                        Total Assets: ${this.props.records.asset_total.toFixed(2)}
                    </Typography>
                    <Typography style={{ marginLeft: '20px' }} variant='subtitle1' component='span'>
                        Total Liabilities: ${this.props.records.liability_total.toFixed(2)}
                    </Typography>
                </div>
            </div>
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

