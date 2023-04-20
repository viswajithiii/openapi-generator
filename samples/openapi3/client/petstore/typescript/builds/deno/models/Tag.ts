/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http.ts';

/**
* A tag for a pet
*/
export class Tag {
    'id'?: number;
    'name'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "number  ",
            "format": "int64"
        },
        {
            "name": "name",
            "baseName": "name",
            "type": "string  ",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Tag.attributeTypeMap;
    }

    public constructor() {

    }
}

