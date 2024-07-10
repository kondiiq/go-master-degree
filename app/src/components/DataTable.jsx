import { DataTable} from 'primereact/datatable';
import { Column } from 'primereact/column';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import 'primereact/resources/primereact.min.css';
import React from 'react';
import 'primereact/resources/themes/lara-light-indigo/theme.css';
import { useState } from 'react';
import { FilterMatchMode } from 'primereact/api';
import { InputText } from 'primereact/inputtext';

const DataTableComponent = () => {
    const [oFilters, setFilters] = useState({
        global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    });
    const oData = [
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
    ];

    return (
        <div className="App">
            <InputText placeholder='Enter name ' onInput={(oEvent) => setFilters({
                ...oFilters,
                global: { value: oEvent.target.value, matchMode: FilterMatchMode.CONTAINS },
            })}/>
            <DataTable  sortField="name" sortOrder={-1} value={oData} sortMode="multiple" filters={oFilters} 
            paginator rows={4} rowsPerPageOptions={[1,2,3,4]} totalRecords={oData.length}
            tableStyle={{ minWidth: '60rem' }} showGridlines stripedRows scrollable scrollHeight='500px'>
                <Column field="id" header="ID" hidden/>
                <Column field="name" header="Name" sortable/>
                <Column field="age" header="Age" sortable/>
                <Column field="city" header="City" sortable/>
            </DataTable>
        </div>
    );
}

export default DataTableComponent;