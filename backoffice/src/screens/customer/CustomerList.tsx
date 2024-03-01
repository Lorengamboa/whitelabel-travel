import { isEmpty } from 'lodash';
import { useNavigate } from 'react-router-dom';

import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { Button } from '@mui/material';

import { useDeleteCustomerMutation } from '@/services/queries/customer.query';

function CustomerList(props) {
  const { data } = props;
  const navigate = useNavigate();
  const { mutate: deleteCustomer } = useDeleteCustomerMutation();

  const handleDelete = async (e: React.MouseEvent<HTMLButtonElement>, customerId: string) => {
    // Handle delete action here
    e.stopPropagation();
    await deleteCustomer(customerId);
  };

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Id</TableCell>
            <TableCell>Name</TableCell>
            <TableCell>Email</TableCell>
            <TableCell>Address</TableCell>
            <TableCell>Url</TableCell>
            <TableCell></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {!isEmpty(data) && data.map((row) => (
            <TableRow
              key={row.id}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              onClick={() => navigate(`/customers/${row.id}`)}
            >
              <TableCell component="th" scope="row">
                {row.id}
              </TableCell>
              <TableCell>{row.name}</TableCell>
              <TableCell>{row.email}</TableCell>
              <TableCell>{row.address}</TableCell>
              <TableCell>{row.url}</TableCell>
              <TableCell> {/* Add this TableCell */}
                <Button
                  variant="contained"
                  color="secondary"
                  onClick={(e) => { handleDelete(e, row.id); }}>
                  Delete
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default CustomerList;