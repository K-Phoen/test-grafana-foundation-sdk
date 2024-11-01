<?php

namespace Grafana\Foundation\Azuremonitor;

final class SubscriptionsQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\SubscriptionsQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\SubscriptionsQueryBuilder())',
        ];
            if ($input->rawQuery !== null && $input->rawQuery !== "") {
    
        
    $buffer = 'rawQuery(';
        $arg0 =\var_export($input->rawQuery, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}
