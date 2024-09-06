// 0
// expense_ratio: 0.3,

// 1
// expense_ratio: 0,

// 2
// expense_ration: 0

// 3
// expense_ratio: 0.82,

// 4
// expense_ratio: 0.64,

// 5
// expense_ratio: 0.58,

// 6
// expense_ratio: 0

// 7
// expense_ratio: 0.18,

// 8
// expense_ratio: 0

// 9
// expense_ratio: 0

// 10
// expense_ratio: 0.77,

// 11
// expense_ratio: 0.9,

// 12
// expense_ratio: 1.18,

// 13
// expense_ratio: 0,

// asset_category
// maturation_date


const assets: { total: number, accountName: string, holdingER: number, skip: boolean }[] = [
    // 0
    { accountName: "Account1", total: 10341.01, holdingER: 0, skip: false },
    { accountName: "Account1", total: 11979.70, holdingER: 1.18, skip: false },
    { accountName: "Account1", total: 1092.52, holdingER: 0, skip: false },
    { accountName: "Account1", total: 3683.72, holdingER: 0.82, skip: false },
    { accountName: "Account1", total: 1913.09, holdingER: 0.18, skip: false },
    { accountName: "Account1", total: 3891.44, holdingER: 0.58, skip: false },
    { accountName: "Account1", total: 3626.10, holdingER: 0, skip: true },

    // 1
    { accountName: "Account2", total: 9371.45, holdingER: 0.82, skip: false },
    { accountName: "Account2", total: 2271.85, holdingER: 0, skip: false },
    { accountName: "Account2", total: 12096.14, holdingER: 0, skip: false },
    { accountName: "Account2", total: 7020.39, holdingER: 0.77, skip: false },
    { accountName: "Account2", total: 13431.37, holdingER: 0, skip: false },
    { accountName: "Account2", total: 7728.62, holdingER: 0.64, skip: false },
    { accountName: "Account2", total: 11658.71, holdingER: 0.3, skip: true },

    // 2
    { accountName: "Account3", total: 7780.79, holdingER: 0.64, skip: false },
    { accountName: "Account3", total: 6969.85, holdingER: 0, skip: false },
    { accountName: "Account3", total: 10443.17, holdingER: 0.9, skip: false },
    { accountName: "Account3", total: 12502.99, holdingER: 0.58, skip: false },
    { accountName: "Account3", total: 286.65, holdingER: 0, skip: false },
    { accountName: "Account3", total: 15706.22, holdingER: 0.82, skip: false },
    { accountName: "Account3", total: 5880.80, holdingER: 0, skip: true },

    // 3
    { accountName: "Account4", total: 7625.72, holdingER: 0.82, skip: false },
    { accountName: "Account4", total: 10586.65, holdingER: 0.18, skip: false },
    { accountName: "Account4", total: 7324.53, holdingER: 0.64, skip: false },
    { accountName: "Account4", total: 1168.83, holdingER: 0, skip: false },
    { accountName: "Account4", total: 5268.47, holdingER: 0.3, skip: false },
    { accountName: "Account4", total: 4929.89, holdingER: 0, skip: false },
    { accountName: "Account4", total: 437.78, holdingER: 0.9, skip: true },

    // 4
    { accountName: "Account5", total: 15407.15, holdingER: 0, skip: false },
    { accountName: "Account5", total: 4662.95, holdingER: 0, skip: false },
    { accountName: "Account5", total: 13466.88, holdingER: 0.64, skip: false },
    { accountName: "Account5", total: 4747.47, holdingER: 1.18, skip: false },
    { accountName: "Account5", total: 515.56, holdingER: 0.77, skip: false },
    { accountName: "Account5", total: 6889.24, holdingER: 0, skip: false },
    { accountName: "Account5", total: 13695.74, holdingER: 0.64, skip: true },

    // 5
    { accountName: "Account6", total: 12673.13, holdingER: 0.64, skip: false },
    { accountName: "Account6", total: 3073.30, holdingER: 1.18, skip: false },
    { accountName: "Account6", total: 11547.25, holdingER: 0, skip: false },
    { accountName: "Account6", total: 2240.44, holdingER: 0, skip: false },
    { accountName: "Account6", total: 3483.74, holdingER: 0, skip: false },
    { accountName: "Account6", total: 1858.15, holdingER: 0.18, skip: false },
    { accountName: "Account6", total: 9448.52, holdingER: 0, skip: true },

    // 6
    { accountName: "Account7", total: 15130.08, holdingER: 0, skip: false },
    { accountName: "Account7", total: 11453.51, holdingER: 0, skip: false },
    { accountName: "Account7", total: 7429.49, holdingER: 0.9, skip: false },
    { accountName: "Account7", total: 1639.83, holdingER: 0.82, skip: false },
    { accountName: "Account7", total: 14835.84, holdingER: 0.58, skip: false },
    { accountName: "Account7", total: 9678.95, holdingER: 1.18, skip: false },
    { accountName: "Account7", total: 10854.47, holdingER: 0.77, skip: true },

    // 7
    { accountName: "Account8", total: 15737.92, holdingER: 0, skip: false },
    { accountName: "Account8", total: 9434.05, holdingER: 1.18, skip: false },
    { accountName: "Account8", total: 10096.92, holdingER: 0.58, skip: false },
    { accountName: "Account8", total: 15411.10, holdingER: 0, skip: false },
    { accountName: "Account8", total: 6899.97, holdingER: 0, skip: false },
    { accountName: "Account8", total: 4343.31, holdingER: 0.9, skip: false },
    { accountName: "Account8", total: 4277.00, holdingER: 0.64, skip: true },

    // 8
    { accountName: "Account9", total: 14045.3, holdingER: 0, skip: false },
    { accountName: "Account9", total: 11642.63, holdingER: 0, skip: false },
    { accountName: "Account9", total: 6618.72, holdingER: 0, skip: false },
    { accountName: "Account9", total: 1124.46, holdingER: 0, skip: false },
    { accountName: "Account9", total: 13441.36, holdingER: 0.58, skip: false },
    { accountName: "Account9", total: 10426.11, holdingER: 0.3, skip: false },
    { accountName: "Account9", total: 11859.11, holdingER: 0, skip: true },
]

console.log(getSnapshotTotals());
console.log(getSnapshotByAccountSummary());

function getSnapshotTotals(): { sum: number, er: number } {
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