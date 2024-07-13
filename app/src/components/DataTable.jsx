import { DataTable} from 'primereact/datatable';
import { Column } from 'primereact/column';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import 'primereact/resources/primereact.min.css';
import React, { useState, useRef } from 'react';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import { FilterMatchMode, PrimeIcons } from 'primereact/api';
import { InputText } from 'primereact/inputtext';
import { Button } from 'primereact/button';
import { Toast } from 'primereact/toast';
import 'primeicons/primeicons.css';
import { ConfirmDialog, confirmDialog } from 'primereact/confirmdialog';
import { Messages } from 'primereact/messages';

const DataTableComponent = () => {
    const [oFilters, setFilters] = useState({
        global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    });
    const [aData, setAData] = useState([
        {
            id : 1,
            name : "Josh",
            age : 25,
            city : "Tbg"
        },
        {
            id : 2,
            name : "Pajeet",
            age : 69,
            city : "London"
        },
        {
            id : 3,
            name : "Ashok",
            age : 22,
            city : "Bengalure"
        },
        {
            id : 4,
            name : "Kumar",
            age : 23,
            city : "XDD"
        }
    ]);

    const fnAddEmptyRow = () => {
        const oNewRow = {
            id : aData.length + 1,
            name : '',
            age : null,
            city : ''
        };
        setAData([...aData, oNewRow]);
    };

    const fnRemoveRow = (oRowData) => {
        confirmDialog({
            message: 'Are you sure you want to delete this record?',
            header: 'Confirm',
            icon: 'pi pi-exclamation-triangle',
            accept: () => {
                setAData(aData.filter(oRow => oRow.id !== oRowData.id));
                fnShowSuccess();
            },
            reject: () => {fnShowErrMessage()}
        });
    };

    const fnEsditRow = (oRowData) => {
        const oReplacedRow = {
            name : '',
            age : null,
            city : ''
        };
        setAData([...aData, oReplacedRow]);
    };

    const oToast = useRef(null);
    const oMessage = useRef(null);

    const fnShowSuccess = () => {
        oToast.current.show({severity: 'success', summary: 'Success Message', detail: 'Message Content', life: 3000});
    };

    const fnShowError = () => {
        oToast.current.show({ severity: 'error', summary: 'Error Message', detail: 'Message Content', life: 3000 });
    };

    const fnShowErrMessage = () => {
        oMessage.current.show( {sticky: true, severity: 'info', summary: 'Info', detail: 'You cancelled delete row', closable: false});
    }
    

    const actionBodyTemplate = (oRowData) => {
        return (
            <React.Fragment>
                <Button className='delete-row' icon={PrimeIcons.MINUS} onClick={() => fnRemoveRow(oRowData)}/>
                <Button className='edit-row' icon={PrimeIcons.USER_EDIT} onClick={() => fnEsditRow(oRowData)}/>
            </React.Fragment>
        );
    };

    return (
        <div className="App">
            <Messages ref={oMessage}/>
            <Toast ref={oToast}/>
            <ConfirmDialog/>
            <InputText
                style={{ display: 'flex', justifyContent: 'flex-end' }}
                placeholder='Enter name '
                onInput={(oEvent) => setFilters({
                    ...oFilters,
                    global: { value: oEvent.target.value, matchMode: FilterMatchMode.CONTAINS },
                })}
            />
            <Button className="add-row" icon={PrimeIcons.PLUS} onClick={() => fnAddEmptyRow()}/>
            <DataTable
                id="knapsackTable"
                value={aData}
                dataKey="id"
                selectionMode="single"
                sortField="name"
                sortOrder={-1}
                sortMode="multiple"
                filters={oFilters}
                paginator
                rows={aData.length}
                rowsPerPageOptions={[25, 50, 100]}
                totalRecords={aData.length}
                tableStyle={{ minWidth: '60rem' }}
                showGridlines
                stripedRows
                scrollable
                scrollHeight='500px'>
                <Column field="id" header="ID" hidden/>
                <Column field="name" header="Name" sortable/>
                <Column field="age" header="Age" sortable/>
                <Column field="city" header="City" sortable/>
                <Column body={actionBodyTemplate} header="Actions"/>
            </DataTable>
        </div>
    );
}

export default DataTableComponent;