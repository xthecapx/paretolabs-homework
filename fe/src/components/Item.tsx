import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import { forwardRef } from "react";
import { useNavigate } from "react-router-dom";

interface ItemProps {
  text: string;
  to: string;
}

export const Item = forwardRef<HTMLLIElement, ItemProps>(
  ({ text, to }, ref) => {
    const navigate = useNavigate();
    const goToPage = () => {
      navigate(to);
    };

    return ref ? (
      <ListItem onClick={goToPage} ref={ref} disablePadding>
        <ListItemButton>
          <ListItemText primary={text} />
        </ListItemButton>
      </ListItem>
    ) : (
      <ListItem onClick={goToPage} disablePadding>
        <ListItemButton>
          <ListItemText primary={text} />
        </ListItemButton>
      </ListItem>
    );
  }
);
