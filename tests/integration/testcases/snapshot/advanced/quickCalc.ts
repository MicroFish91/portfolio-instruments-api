// 0
//     asset_category: "TSM",

// 1
//     asset_category:  "LTB",

// 2
//     asset_category: "CASH",

// 3
//     asset_category: "GOLD",

// 4
//     asset_category: "DSCV",

// 5
//     asset_category: "OTHER",

// 6
//     asset_category:  "STB",

// 7
//     asset_category: "TSM",

// 8
//     asset_category:  "LTB",

// 9
//     asset_category: "CASH",

// 10
//     asset_category: "GOLD",

// 11
//     asset_category: "DSCV",

// 12
//     asset_category: "OTHER",

// 13
//     asset_category:  "ITB",

// asset_category
// maturation_date

// Equivalent to the assets provided under `createsnapshot.go`
const assets: { total: number, assetCategory: string, accountName: string, holdingER: number, skip: boolean }[] = [
    // 0
    // Holdings: 2, 12, 6, 3, 7, 5, 9
    { assetCategory: "CASH", accountName: "Account1", total: 10341.01, holdingER: 0, skip: false },
    { assetCategory: "OTHER", accountName: "Account1", total: 11979.70, holdingER: 1.18, skip: false },
    { assetCategory: "STB", accountName: "Account1", total: 1092.52, holdingER: 0, skip: false },
    { assetCategory: "GOLD", accountName: "Account1", total: 3683.72, holdingER: 0.82, skip: false },
    { assetCategory: "TSM", accountName: "Account1", total: 1913.09, holdingER: 0.18, skip: false },
    { assetCategory: "OTHER", accountName: "Account1", total: 3891.44, holdingER: 0.58, skip: false },
    { assetCategory: "CASH", accountName: "Account1", total: 3626.10, holdingER: 0, skip: true },

    // 1
    // Holdings: 3, 1, 13, 10, 8, 4, 0
    { assetCategory: "GOLD", accountName: "Account2", total: 9371.45, holdingER: 0.82, skip: false },
    { assetCategory: "LTB", accountName: "Account2", total: 2271.85, holdingER: 0, skip: false },
    { assetCategory: "ITB", accountName: "Account2", total: 12096.14, holdingER: 0, skip: false },
    { assetCategory: "GOLD", accountName: "Account2", total: 7020.39, holdingER: 0.77, skip: false },
    { assetCategory: "LTB", accountName: "Account2", total: 13431.37, holdingER: 0, skip: false },
    { assetCategory: "DSCV", accountName: "Account2", total: 7728.62, holdingER: 0.64, skip: false },
    { assetCategory: "TSM", accountName: "Account2", total: 11658.71, holdingER: 0.3, skip: true },

    // 2
    // Holdings: 4, 9, 11, 5, 2, 3, 6
    { assetCategory: "DSCV", accountName: "Account3", total: 7780.79, holdingER: 0.64, skip: false },
    { assetCategory: "CASH", accountName: "Account3", total: 6969.85, holdingER: 0, skip: false },
    { assetCategory: "DSCV", accountName: "Account3", total: 10443.17, holdingER: 0.9, skip: false },
    { assetCategory: "OTHER", accountName: "Account3", total: 12502.99, holdingER: 0.58, skip: false },
    { assetCategory: "CASH", accountName: "Account3", total: 286.65, holdingER: 0, skip: false },
    { assetCategory: "GOLD", accountName: "Account3", total: 15706.22, holdingER: 0.82, skip: false },
    { assetCategory: "STB", accountName: "Account3", total: 5880.80, holdingER: 0, skip: true },

    // 3
    // Holdings: 3, 7, 4, 9, 0, 2, 11
    { assetCategory: "GOLD", accountName: "Account4", total: 7625.72, holdingER: 0.82, skip: false },
    { assetCategory: "TSM", accountName: "Account4", total: 10586.65, holdingER: 0.18, skip: false },
    { assetCategory: "DSCV", accountName: "Account4", total: 7324.53, holdingER: 0.64, skip: false },
    { assetCategory: "CASH", accountName: "Account4", total: 1168.83, holdingER: 0, skip: false },
    { assetCategory: "TSM", accountName: "Account4", total: 5268.47, holdingER: 0.3, skip: false },
    { assetCategory: "CASH", accountName: "Account4", total: 4929.89, holdingER: 0, skip: false },
    { assetCategory: "DSCV", accountName: "Account4", total: 437.78, holdingER: 0.9, skip: true },

    // 4
    // Holdings: 8, 9, 4, 12, 10, 6, 4
    { assetCategory: "LTB", accountName: "Account5", total: 15407.15, holdingER: 0, skip: false },
    { assetCategory: "CASH", accountName: "Account5", total: 4662.95, holdingER: 0, skip: false },
    { assetCategory: "DSCV", accountName: "Account5", total: 13466.88, holdingER: 0.64, skip: false },
    { assetCategory: "OTHER", accountName: "Account5", total: 4747.47, holdingER: 1.18, skip: false },
    { assetCategory: "GOLD", accountName: "Account5", total: 515.56, holdingER: 0.77, skip: false },
    { assetCategory: "STB", accountName: "Account5", total: 6889.24, holdingER: 0, skip: false },
    { assetCategory: "DSCV", accountName: "Account5", total: 13695.74, holdingER: 0.64, skip: true },

    // 5
    // Holdings: 4, 12, 2, 6, 1, 7, 13
    { assetCategory: "DSCV", accountName: "Account6", total: 12673.13, holdingER: 0.64, skip: false },
    { assetCategory: "OTHER", accountName: "Account6", total: 3073.30, holdingER: 1.18, skip: false },
    { assetCategory: "CASH", accountName: "Account6", total: 11547.25, holdingER: 0, skip: false },
    { assetCategory: "STB", accountName: "Account6", total: 2240.44, holdingER: 0, skip: false },
    { assetCategory: "LTB", accountName: "Account6", total: 3483.74, holdingER: 0, skip: false },
    { assetCategory: "TSM", accountName: "Account6", total: 1858.15, holdingER: 0.18, skip: false },
    { assetCategory: "ITB", accountName: "Account6", total: 9448.52, holdingER: 0, skip: true },

    // 6
    // Holdings: 1, 13, 11, 3, 5, 12, 10
    { assetCategory: "LTB", accountName: "Account7", total: 15130.08, holdingER: 0, skip: false },
    { assetCategory: "ITB", accountName: "Account7", total: 11453.51, holdingER: 0, skip: false },
    { assetCategory: "DSCV", accountName: "Account7", total: 7429.49, holdingER: 0.9, skip: false },
    { assetCategory: "GOLD", accountName: "Account7", total: 1639.83, holdingER: 0.82, skip: false },
    { assetCategory: "OTHER", accountName: "Account7", total: 14835.84, holdingER: 0.58, skip: false },
    { assetCategory: "OTHER", accountName: "Account7", total: 9678.95, holdingER: 1.18, skip: false },
    { assetCategory: "GOLD", accountName: "Account7", total: 10854.47, holdingER: 0.77, skip: true },

    // 7
    // Holdings: 6, 12, 5, 13, 2, 11, 4
    { assetCategory: "STB", accountName: "Account8", total: 15737.92, holdingER: 0, skip: false },
    { assetCategory: "OTHER", accountName: "Account8", total: 9434.05, holdingER: 1.18, skip: false },
    { assetCategory: "OTHER", accountName: "Account8", total: 10096.92, holdingER: 0.58, skip: false },
    { assetCategory: "ITB", accountName: "Account8", total: 15411.10, holdingER: 0, skip: false },
    { assetCategory: "CASH", accountName: "Account8", total: 6899.97, holdingER: 0, skip: false },
    { assetCategory: "DSCV", accountName: "Account8", total: 4343.31, holdingER: 0.9, skip: false },
    { assetCategory: "DSCV", accountName: "Account8", total: 4277.00, holdingER: 0.64, skip: true },

    // 8
    // Holdings: 2, 9, 6, 8, 5, 0, 13
    { assetCategory: "CASH", accountName: "Account9", total: 14045.3, holdingER: 0, skip: false },
    { assetCategory: "CASH", accountName: "Account9", total: 11642.63, holdingER: 0, skip: false },
    { assetCategory: "STB", accountName: "Account9", total: 6618.72, holdingER: 0, skip: false },
    { assetCategory: "LTB", accountName: "Account9", total: 1124.46, holdingER: 0, skip: false },
    { assetCategory: "OTHER", accountName: "Account9", total: 13441.36, holdingER: 0.58, skip: false },
    { assetCategory: "TSM", accountName: "Account9", total: 10426.11, holdingER: 0.3, skip: false },
    { assetCategory: "ITB", accountName: "Account9", total: 11859.11, holdingER: 0, skip: true },
]

// Back of the envelope calculations for computing the different `Advanced Snapshot` test values
console.log(getSnapshotTotalSummary());
console.log(getSnapshotByAccountSummary());
console.log(getSnapshotByAssetCategorySummary());

function getSnapshotTotalSummary(): { sum: number, er: number } {
    let sum = 0;
    let erSum = 0;
    for (const { total, holdingER } of assets) {
        sum += total;
        erSum += total * holdingER;
    }

    return { sum, er: erSum / sum };
}

function getSnapshotByAccountSummary(): { fields: string[], totals: number[] } {
    const resourcesGrouped: Record<string, number> = {};
    for (const { accountName, total } of assets) {
        resourcesGrouped[accountName] = (resourcesGrouped[accountName] ?? 0) + total;
    }

    const fields: string[] = [];
    const totals: number[] = [];
    for (const [key, val] of Object.entries(resourcesGrouped)) {
        fields.push(key);
        totals.push(val);
    }

    return { fields, totals };
}

function getSnapshotByAssetCategorySummary(): { fields: string[], totals: number[] } {
    const resourcesGrouped: Record<string, number> = {};
    for (const { assetCategory, total } of assets) {
        resourcesGrouped[assetCategory] = (resourcesGrouped[assetCategory] ?? 0) + total;
    }

    const fields: string[] = [];
    const totals: number[] = [];
    for (const [key, val] of Object.entries(resourcesGrouped)) {
        fields.push(key);
        totals.push(val);
    }

    return { fields, totals };
}