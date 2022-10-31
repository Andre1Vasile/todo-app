import styled from "styled-components"
import background from 'C:/Programming/todoapp/frontend/src/media/tile_background.png';

export const ContentSection = styled.div`
        background-image: url(${background});
        padding: 10px 20px;
        border: 2px solid #dddddd;
        border-radius: 3px;
        margin-bottom: 20px;
        box-shadow:  5px 5px 17px #cccccc,
        -5px -5px 17px #ffffff;

        li{
                color:#fafafa,
                border-radius: 3px,
                border-style: round;
        }
`