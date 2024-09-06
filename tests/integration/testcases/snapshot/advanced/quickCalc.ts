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


const assets: { Total: number, holdingER: number, Skip: boolean }[] = [
    // 0
    { Total: 10341.01, holdingER: 0, Skip: false },
    { Total: 11979.70, holdingER: 1.18, Skip: false },
    { Total: 1092.52, holdingER: 0, Skip: false },
    { Total: 3683.72, holdingER: 0.82, Skip: false },
    { Total: 1913.09, holdingER: 0.18, Skip: false },
    { Total: 3891.44, holdingER: 0.58, Skip: false },
    { Total: 3626.10, holdingER: 0, Skip: true },

    // 1
    { Total: 9371.45, holdingER: 0.82, Skip: false },
    { Total: 2271.85, holdingER: 0, Skip: false },
    { Total: 12096.14, holdingER: 0, Skip: false },
    { Total: 7020.39, holdingER: 0.77, Skip: false },
    { Total: 13431.37, holdingER: 0, Skip: false },
    { Total: 7728.62, holdingER: 0.64, Skip: false },
    { Total: 11658.71, holdingER: 0.3, Skip: true },

    // 2
    { Total: 7780.79, holdingER: 0.64, Skip: false },
    { Total: 6969.85, holdingER: 0, Skip: false },
    { Total: 10443.17, holdingER: 0.9, Skip: false },
    { Total: 12502.99, holdingER: 0.58, Skip: false },
    { Total: 286.65, holdingER: 0, Skip: false },
    { Total: 15706.22, holdingER: 0.82, Skip: false },
    { Total: 5880.80, holdingER: 0, Skip: true },

    // 3
    { Total: 7625.72, holdingER: 0.82, Skip: false },
    { Total: 10586.65, holdingER: 0.18, Skip: false },
    { Total: 7324.53, holdingER: 0.64, Skip: false },
    { Total: 1168.83, holdingER: 0, Skip: false },
    { Total: 5268.47, holdingER: 0.3, Skip: false },
    { Total: 4929.89, holdingER: 0, Skip: false },
    { Total: 437.78, holdingER: 0.9, Skip: true },

    // 4
    { Total: 15407.15, holdingER: 0, Skip: false },
    { Total: 4662.95, holdingER: 0, Skip: false },
    { Total: 13466.88, holdingER: 0.64, Skip: false },
    { Total: 4747.47, holdingER: 1.18, Skip: false },
    { Total: 515.56, holdingER: 0.77, Skip: false },
    { Total: 6889.24, holdingER: 0, Skip: false },
    { Total: 13695.74, holdingER: 0.64, Skip: true },

    // 5
    { Total: 12673.13, holdingER: 0.64, Skip: false },
    { Total: 3073.30, holdingER: 1.18, Skip: false },
    { Total: 11547.25, holdingER: 0, Skip: false },
    { Total: 2240.44, holdingER: 0, Skip: false },
    { Total: 3483.74, holdingER: 0, Skip: false },
    { Total: 1858.15, holdingER: 0.18, Skip: false },
    { Total: 9448.52, holdingER: 0, Skip: true },

    // 6
    { Total: 15130.08, holdingER: 0, Skip: false },
    { Total: 11453.51, holdingER: 0, Skip: false },
    { Total: 7429.49, holdingER: 0.9, Skip: false },
    { Total: 1639.83, holdingER: 0.82, Skip: false },
    { Total: 14835.84, holdingER: 0.58, Skip: false },
    { Total: 9678.95, holdingER: 1.18, Skip: false },
    { Total: 10854.47, holdingER: 0.77, Skip: true },

    // 7
    { Total: 15737.92, holdingER: 0, Skip: false },
    { Total: 9434.05, holdingER: 1.18, Skip: false },
    { Total: 10096.92, holdingER: 0.58, Skip: false },
    { Total: 15411.10, holdingER: 0, Skip: false },
    { Total: 6899.97, holdingER: 0, Skip: false },
    { Total: 4343.31, holdingER: 0.9, Skip: false },
    { Total: 4277.00, holdingER: 0.64, Skip: true },

    // 8
    { Total: 14045.3, holdingER: 0, Skip: false },
    { Total: 11642.63, holdingER: 0, Skip: false },
    { Total: 6618.72, holdingER: 0, Skip: false },
    { Total: 1124.46, holdingER: 0, Skip: false },
    { Total: 13441.36, holdingER: 0.58, Skip: false },
    { Total: 10426.11, holdingER: 0.3, Skip: false },
    { Total: 11859.11, holdingER: 0, Skip: true },
]

let sum = 0;
let erSum = 0;
for (const { Total, holdingER, Skip } of assets) {
    sum += Total
    erSum += Total * holdingER
}

console.log("sum:", sum)
console.log("er:", erSum / sum)