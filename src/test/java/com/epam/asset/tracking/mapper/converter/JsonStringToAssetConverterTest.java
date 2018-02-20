package com.epam.asset.tracking.mapper.converter;

import com.epam.asset.tracking.domain.Asset;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class JsonStringToAssetConverterTest {

    @Test
    public void convertJSONStringToAsset(){

        String jsonAsset = "{\"uuid\": \"9d40ee4e-bf1e-4f74-8237-c5e9b6e8f6d3\", \"serialNumber\": \"3VW1W21KIBM312176\", \"assetType\": \"VEHICLE\", \"ownerName\": \"Jhonn Doe\", \"description\": \"2011 VW JETTA STYLE ACTIVE MANUAL TRANSMISION. SIDE AIRBAGS PACKAGE, COLOR: WHITE CANDY INTERIOR COLOR: BLACK FABRIC . ENGINE: 2.5L FIVE CYLINDERS ENGINE NUMBER: CCC094323 MADE IN: MEXICO BUYER NAME: JHONN DOE ADDRESS: 123 ABBY ROAD, THE DOMAIN. AUTIN TEXAS, USA. SELLER NAME: RAY REDDINGTON\"}";

        JsonStringToAssetConverter converter = new JsonStringToAssetConverter();


        Asset asset = null;

        //Converts the given JSON Script into an Asset Object
        asset  = converter.convert(jsonAsset, null);

        //Check if the values from the Asset object created matched the on the JSON String
        assertEquals(asset.getUuid(), "9d40ee4e-bf1e-4f74-8237-c5e9b6e8f6d3");
        assertEquals(asset.getSerialNumber(), "3VW1W21KIBM312176");
        assertEquals(asset.getAssetType(), "VEHICLE");
        assertEquals(asset.getDescription(),"2011 VW JETTA STYLE ACTIVE MANUAL TRANSMISION. SIDE AIRBAGS PACKAGE, COLOR: WHITE CANDY INTERIOR COLOR: BLACK FABRIC . ENGINE: 2.5L FIVE CYLINDERS ENGINE NUMBER: CCC094323 MADE IN: MEXICO BUYER NAME: JHONN DOE ADDRESS: 123 ABBY ROAD, THE DOMAIN. AUTIN TEXAS, USA. SELLER NAME: RAY REDDINGTON");


    }
}